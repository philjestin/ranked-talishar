"use client";

import React, { FormEvent } from "react";
import { useRouter } from "next/router";

export default function Login() {
  // const router = useRouter();

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const username = formData.get("username");
    const password = formData.get("password");

    const response = await fetch("http://localhost:8000/api/users/login", {
      method: "POST",
      headers: { "Content-Type": "text/plain" },
      body: JSON.stringify({ username: username, password }),
    });

    if (response.ok) {
      console.log({ response })
      const data = await response.json();
      console.log({ data })
      // router.push("/matches");
    } else {
      console.log({ response })
      // Handle errors
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="username" placeholder="username" required />
      <input type="password" name="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  );
}
