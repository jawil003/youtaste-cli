export interface Poll {
  restaurantName: string;
  provider?: "youtaste" | "lieferando";
  count?: number;
}
