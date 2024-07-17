import { Form, useLoaderData } from "@remix-run/react";
import type { FunctionComponent } from "react";
import { getContacts } from "~/client/contacts/client-api/getContacts";


export const loader = async () => {
  const contacts = await getContacts();

  return contacts;
};

export default function Contact() {
  const data = useLoaderData<typeof loader>();
  
  console.log({ data })

  return (
    <div id="contact">
      Contacts id Page
    </div>
  )
}