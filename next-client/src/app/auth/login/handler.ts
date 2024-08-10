import { NextApiRequest, NextApiResponse } from "next";
import { signIn } from "@/auth";
import { NextResponse } from "next/server";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method === 'OPTIONS') {
    console.log("what the fuck")
    return res.status(200).send('ok');
  }


  try {
    const { email, password } = req.body;
    await signIn("credentials", { email, password });

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
