import * as grpcWeb from "grpc-web";

export interface GRPCClientResponse<T> {
  error: grpcWeb.RpcError | null;
  response: T | null;
}

export type ErrorHandler = (
  error: grpcWeb.RpcError,
  serviceName?: string
) => void;

export interface GRPCClientConfig {
  hostname: string;
  serviceName?: string;
  onError?: ErrorHandler;
}
