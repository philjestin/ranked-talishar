import { getContacts } from "../client-api/getContacts";

export const loader = async () => {
  const contacts = await getContacts();

  return contacts
}