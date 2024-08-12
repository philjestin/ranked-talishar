import { NextApiRequest, NextApiResponse } from "next";
// import { signIn } from "@/auth";
import { NextRequest, NextResponse } from "next/server";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  console.log("LOGIN>TS")
  if (req.method === 'OPTIONS') {
    console.log("what the fuck")
    return res.status(200).send('ok');
  }


  try {
    const { username, password } = req.body;
    const response = await fetch("http://localhost:8000/api/users/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username: username, password }),
    });

      const data = await response.json();
      console.log({ data })

    res.status(200).json({ success: true });
  } catch (error) {
    if (error.type === "CredentialsSignin") {
      res.status(401).json({ error: "Invalid credentials." });
    } else {
      res.status(500).json({ error: "Something went wrong." });
    }
  }
}

export const OPTIONS = async (request: NextRequest) => {
  return new NextResponse('', {
    status: 200
  })
}
