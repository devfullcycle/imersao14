import type { DirectionsResponseData } from "@googlemaps/google-maps-services-js";

export type Route = {
  id: string;
  name: string;
  source: { name: string; location: { lat: number; lng: number } };
  destination: { name: string; location: { lat: number; lng: number } };
  distance: number;
  duration: number;
  directions: DirectionsResponseData & { request: any };
  created_at: Date;
  updated_at: Date;
};
