export interface RegisterUserProps {
  username: string
  password: string
  email: string
}

interface LoginUserProps {
  password: string;
  user_name: string;
}

const baseUrl = "http://localhost:8000/api/users";

export async function registerUserService(userData: RegisterUserProps) {
  const url = `${baseUrl}/signup`;
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
    const data = await response.json();
    console.log({ data });
    console.log(data.body)
    return data;
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

