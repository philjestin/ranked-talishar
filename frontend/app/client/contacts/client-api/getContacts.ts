export async function getContacts () {
  const data = await fetch("localhost:8000/api/contacts")
  const contacts = await data.json();

  return contacts
}