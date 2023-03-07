export enum StatusEnum {
  "IDLE" = "IDLE",
  "LOADING" = "LOADING",
  "SUCCEEDED" = "SUCCEEDED",
  "FAILED" = "FAILED",
}

export const isLoading = (status: StatusEnum) => status === StatusEnum.LOADING;
