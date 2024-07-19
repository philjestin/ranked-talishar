export interface RegisterUserProps {
  user_email: string
  password: string
  user_name: string
}

interface LoginUserProps {
  password: string;
  user_name: string;
}

const baseUrl = "http://localhost:8000/api/users";

export async function registerUserService(userData: RegisterUserProps) {
  const url = `${baseUrl}`;
  console.log({ userData });
  console.log(JSON.stringify({...userData}));
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "text/plain",
      },
      body: JSON.stringify({ ...userData }),
      cache: "no-cache",
    });
    console.log({ response })
    return response.json();
  } catch (error) {
    console.error("Error registering user: ", error);
  }
}

export async function loginUserService(userData: LoginUserProps) {
  const url = new URL(`${baseUrl}/login`);

  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...userData }),
      cache: "no-cache",
    });

    return response.json();
  } catch (error) {
    console.error("Login Service Error:", error);
    throw error;
  }
}

