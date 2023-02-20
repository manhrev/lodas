import { StatusCode as grpcStatusCode } from "grpc-web";

import { GRPCClientResponse, ErrorHandler, GRPCClientConfig } from "./types";

export const KEY_ACCESS_TOKEN = "accessToken";
export const QUERY_PARAM_ACCESS_TOKEN = "token";
export const IGNORE_TOAST_CODES = [grpcStatusCode.UNAUTHENTICATED];

class gRPCClientAbstract {
  client: any = null;
  clientName: string = "";
  serviceName: string;
  onError?: ErrorHandler;

  constructor(Client: any, config: GRPCClientConfig) {
    this.clientName = Client?.name || "NoClientName";
    this.client = new Client(config.hostname);
    this.onError = config.onError;
    this.serviceName = config.serviceName || "UNNAMED";
  }

  logFuncName(funcName: string) {
    return `${this.clientName}.${funcName}`;
  }

  async gRPCClientRequest<T>(
    func: string,
    request: any,
    option: any = {}
  ): Promise<GRPCClientResponse<T>> {
    try {
      const token = localStorage.getItem(KEY_ACCESS_TOKEN);
      if (token) {
        option = { ...option, Authorization: `Bearer ${token}` };
      }

      console.log(
        `%c gRPCClientRequest -> [${this.logFuncName(func)}] -> REQUEST:`,
        "background-color: #deeb34; color: #000; font-size: 14px"
      );
      console.log(">>> request:", request.toObject());
      console.log(">>> option:", option);

      const response = await this.client[func](request, option);

      console.log(
        `%c>>>>> gRPCClientResponse -> [${this.logFuncName(func)}] -> SUCCESS:`,
        "background-color: #23d947; color: #000; font-size: 14px",
        response.toObject()
      );

      return {
        error: null,
        response: response.toObject(),
      };
    } catch (error: any) {
      switch (error?.code) {
        case grpcStatusCode.UNAUTHENTICATED:
          console.log(
            `%c>>>>> gRPCClientResponse -> [${this.logFuncName(
              func
            )}] -> ERROR -> UNAUTHENTICATED: `,
            "background-color: #c0392b; color: #000; font-size: 14px",
            error
          );
          break;

        case grpcStatusCode.UNKNOWN:
        case grpcStatusCode.UNIMPLEMENTED:
        case grpcStatusCode.INVALID_ARGUMENT:
        case grpcStatusCode.NOT_FOUND:
          console.log(
            `%c>>>>> gRPCClientResponse -> [${this.logFuncName(
              func
            )}] -> ERROR: `,
            "background-color: #c0392b; color: #000; font-size: 14px",
            error
          );
          break;

        default:
          console.log(
            `%c>>>>> gRPCClientResponse  -> [${this.logFuncName(
              func
            )}] -> ERROR: `,
            "background-color: #c0392b; color: #000; font-size: 14px",
            error
          );
          break;
      }

      this.onError && this.onError(error, this.serviceName);

      return {
        error,
        response: null,
      };
    }
  }
}

export default gRPCClientAbstract;
