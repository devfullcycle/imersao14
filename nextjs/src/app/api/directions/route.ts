import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const url = new URL(request.url);
  const originId = url.searchParams.get("originId");
  const destinationId = url.searchParams.get("destinationId");
  const response = await fetch(`http://host.docker.internal:3000/directions?originId=${originId}&destinationId=${destinationId}`, {
    next: {
      revalidate: 60, //cache bem grande, valor bem grande
    },
  });
  return NextResponse.json(await response.json());
}
