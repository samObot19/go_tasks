package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book *models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []*models.Book
	ListBorrowedBooks(memberID int) []*models.Book
	AddMember(member *models.Member)
	RemoveMember(memberID int) error
}

type Manager struct {
	id int
	books   map[int]*models.Book
	members map[int]*models.Member
}

func NewLibraryManager() *Manager {
	return &Manager{
		books:   make(map[int]*models.Book),
		members: make(map[int]*models.Member),
	}
}

func (m *Manager) AddBook(book *models.Book) {
	m.books[m.id++] = book
}

func (m *Manager) RemoveBook(bookID int) error {
	_, ok := m.books[bookID]
	if !ok {
		return errors.New("the book with the given ID does not exist")
	}

	delete(m.books, bookID)
	return nil
}

func (m *Manager) BorrowBook(bookID int, memberID int) error {
	book, b_ok := m.books[bookID]
	member, m_ok := m.members[memberID]

	if !b_ok {
		return errors.New("the book with the given ID does not exist")
	}

	if !m_ok {
		return errors.New("the member with the given ID does not exist")
	}

	if book.Status == "Not available" {
		return errors.New("the book is not available")
	}

	book.Status = "Not available"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (m *Manager) ReturnBook(bookID int, memberID int) error {
	book, b_ok := m.books[bookID]
	member, m_ok := m.members[memberID]

	if !b_ok {
		return errors.New("the book with the given ID does not exist")
	}

	if !m_ok {
		return errors.New("the member with the given ID does not exist")
	}

	if book.Status == "available" {
		return errors.New("the book is not currently borrowed")
	}

	book.Status = "available"

	index := -1
	for ind, b := range member.BorrowedBooks {
		if b.ID == bookID {
			index = ind
			break
		}
	}

	if index == -1 {
		return errors.New("the book is not currently borrowed by the given member")
	}

	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
	return nil
}

func (m *Manager) ListAvailableBooks() []*models.Book {
	var availableBooks []*models.Book
	for _, book := range m.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (m *Manager) ListBorrowedBooks(memberID int) []*models.Book {
	member, ok := m.members[memberID]
	if !ok {
		return nil
	}

	return member.BorrowedBooks
}

func (m *Manager) AddMember(member *models.Member) {
	m.members[member.ID] = member
}

func (m *Manager) RemoveMember(memberID int) error {
	_, ok := m.members[memberID]
	if !ok {
		return errors.New("the member with the given ID does not exist")
	}

	delete(m.members, memberID)
	return nil
}

