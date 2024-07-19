'use server';

import {
  registerUserService,
  RegisterUserProps,
} from "./services/auth-service";

export async function registerUserAction(prevState: any, formData: FormData) {
  const fields = {
    user_name: formData.get("username"),
    password: formData.get("password"),
    user_email: formData.get("email")
  }

  const responseData = await registerUserService(fields as RegisterUserProps);

  return {
    ...prevState,
    data: fields
  }
}