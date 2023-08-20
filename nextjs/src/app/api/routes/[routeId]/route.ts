import { NextRequest, NextResponse } from "next/server";

export async function GET(
  request: NextRequest,
  { params }: { params: { routeId: string } }
) {
  const id = params.routeId;
  const response = await fetch(`http://host.docker.internal:3000/routes/${id}`, {
    next: {
      revalidate: 60,
    },
  });
  return NextResponse.json(await response.json());
}
