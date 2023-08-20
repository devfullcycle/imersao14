"use client";

import { useEffect, useRef } from "react";
import { useMap } from "../hooks/useMap";
import { Route } from "../utils/model";
import { socket } from "../utils/socket-io";

export function AdminPage() {
  const mapContainerRef = useRef<HTMLDivElement>(null);
  const map = useMap(mapContainerRef);

  useEffect(() => {
    socket.connect();

    socket.on(
      "admin-new-points",
      async (data: { route_id: string; lat: number; lng: number }) => {
        console.log(data);
        if (!map?.hasRoute(data.route_id)) {
          const response = await fetch(`${process.env.NEXT_PUBLIC_NEXT_API_URL}/routes/${data.route_id}`);
          const route: Route = await response.json();
          map?.removeRoute(data.route_id);
          await map?.addRouteWithIcons({
            routeId: data.route_id,
            startMarkerOptions: {
              position: route.directions.routes[0].legs[0].start_location,
            },
            endMarkerOptions: {
              position: route.directions.routes[0].legs[0].end_location,
            },
            carMarkerOptions: {
              position: route.directions.routes[0].legs[0].start_location,
            },
          });
        }
        map?.moveCar(data.route_id, {
          lat: data.lat,
          lng: data.lng,
        });
      }
    );

    return () => {
      socket.disconnect();
    };
  }, [map]);

  return (
    <div
      style={{
        height: "100%",
        width: "100%",
      }}
    >
      <div
        id="map"
        style={{
          height: "100%",
          width: "100%",
        }}
        ref={mapContainerRef}
      ></div>
    </div>
  );
}

export default AdminPage;
