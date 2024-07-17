import ContactList from "./ContactList";

async function getContacts() {
  console.log({ fetch })
  const data = await fetch("http://localhost:8000/api/contacts", {
    headers: {
      'Content-Type': 'application/json',
    }
  });

  if (!data.ok) {
    throw new Error("Failed to fetch contacts")
  }
  const contacts = await data.json();

  return contacts;
}

export default async function Contacts() {
  const data = await getContacts();

  return (
    <main>
      <h1>Contacts</h1>
      <ContactList contacts={data.contacts} />
    </main>
  );
}
