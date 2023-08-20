import { revalidateTag } from "next/cache";
import { NextResponse } from "next/server";

export async function GET() {
  const response = await fetch("http://host.docker.internal:3000/routes", {
    next: {
      revalidate: 1, //cache bem grande, valor bem grande,
      tags: ["routes"], //revalidação de cache sob demanda
    },
  });
  return NextResponse.json(await response.json());
}

export async function POST(request: Request) {
  const response = await fetch("http://host.docker.internal:3000/routes", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(await request.json()),
  });
  revalidateTag("routes");
  return NextResponse.json(await response.json());
}
