'use client';
import { useFormState } from "react-dom";

import { registerUserAction } from "@/app/data/actions/auth-actions";

export default function SignUpForm() {
  const initialState = {
    data: null,
  }

  const [formState, formAction] = useFormState(
    registerUserAction,
    initialState
  );


  return (
    <form action={formAction}>
      <input type="email" name="email" placeholder="Email" required />
      <input type="password" name="password" placeholder="Password" required />
      <input type="username" name="username" placeholder="username" required />
      <button type="submit">Login</button>
    </form>
  );
}
