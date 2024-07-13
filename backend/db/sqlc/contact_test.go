package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/philjestin/ranked-talishar/test_util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Contact {
	arg := CreateContactParams{
		FirstName:   test_util.RandomFirstName(),
		LastName:    test_util.RandomLastName(),
		PhoneNumber: test_util.RandomPhoneNumber(),
		Street:      test_util.RandomStreet(),
	}

	contact, err := testQueries.CreateContact(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, contact)

	require.Equal(t, arg.FirstName, contact.FirstName)
	require.Equal(t, arg.LastName, contact.LastName)
	require.Equal(t, arg.PhoneNumber, contact.PhoneNumber)
	require.Equal(t, arg.Street, contact.Street)

	require.NotZero(t, contact.ContactID)
	require.NotZero(t, contact.CreatedAt)

	return contact
}

func TestCreateContact(t *testing.T) {
	createRandomAccount(t)
}

func TestGetContact(t *testing.T) {
	// Create account
	contact1 := createRandomAccount(t)
	contact2, err := testQueries.GetContactById(context.Background(), contact1.ContactID)

	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ContactID, contact2.ContactID)
	require.Equal(t, contact1.FirstName, contact2.FirstName)
	require.Equal(t, contact1.LastName, contact2.LastName)
	require.Equal(t, contact1.PhoneNumber, contact2.PhoneNumber)
	require.Equal(t, contact1.Street, contact2.Street)

	require.WithinDuration(t, contact1.CreatedAt, contact2.CreatedAt, time.Second)
}

func TestUpdateContact(t *testing.T) {
	contact1 := createRandomAccount(t)

	arg := UpdateContactParams{
		ContactID: contact1.ContactID,
		FirstName: sql.NullString{String: test_util.RandomFirstName(), Valid: true},
	}

	contact2, err := testQueries.UpdateContact(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contact2)

	require.Equal(t, contact1.ContactID, contact2.ContactID)
	require.Equal(t, arg.FirstName.String, contact2.FirstName)
	require.Equal(t, contact1.LastName, contact2.LastName)
	require.Equal(t, contact1.PhoneNumber, contact2.PhoneNumber)
	require.Equal(t, contact1.Street, contact2.Street)
	require.WithinDuration(t, contact1.CreatedAt, contact2.CreatedAt, time.Second)

	require.NotZero(t, contact2.UpdatedAt)
}

func TestDeleteContact(t *testing.T) {
	contact1 := createRandomAccount(t)
	err := testQueries.DeleteContact(context.Background(), contact1.ContactID)
	require.NoError(t, err)

	contact2, err := testQueries.GetContactById(context.Background(), contact1.ContactID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, contact2)
}

func TestListContacts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListContactsParams{
		Limit:  5,
		Offset: 5,
	}

	contacts, err := testQueries.ListContacts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, contacts, 5)

	for _, contact := range contacts {
		require.NotEmpty(t, contact)
	}
}
