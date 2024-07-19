'use server';

import {
  registerUserService,
  RegisterUserProps,
} from "./services/auth-service";

export async function registerUserAction(prevState: any, formData: FormData) {
  const fields = {
    username: formData.get("username"),
    password: formData.get("password"),
    email: formData.get("email")
  }

  const responseData = await registerUserService(fields as RegisterUserProps);

  return {
    ...prevState,
    data: fields
  }
}