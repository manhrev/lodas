import { StatusEnum } from "../constant";

export interface CommonState {
  error?: string;
  status: StatusEnum;
}
