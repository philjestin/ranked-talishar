'use client'

interface Props {
  contacts: any[];
}

export default function ContactList(props: Props) {
  const { contacts } = props;

  if (!contacts) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {contacts &&
        contacts.length > 0 &&
        contacts.map((contact: any) => {
          return (
            <div key={`${contact.contact_id}`}>
              {contact.first_name} {contact.last_name}
            </div>
          );
        })}
    </>
  );
}