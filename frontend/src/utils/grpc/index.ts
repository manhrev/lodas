import { StatusCode as grpcStatusCode } from "grpc-web";
import rpcAuthClient from "./client/auth";
import { Code as AuthCode } from "../../lib/auth/auth_code_pb";
import rpcLodasClient from "./client/lodas";

export * from "./abstract/gRPCClient";
export * from "./abstract/types";

const onAuthError = (error: any, serviceName: any) => {
  // switch (error?.code) {
  //   case grpcStatusCode.UNAUTHENTICATED:
  //     onLogout(error.code);
  //     break;
  // }
  // if (IGNORE_TOAST_CODES.includes(error?.code)) return;
  // iToast.error({
  //   title: i18n.t('Common::ERROR_TITLE'),
  //   msg: i18n.t(`Error::${serviceName}.${error?.code}`),
  // });
  // error?.code && alert(error.code);
};

const HOST = `http://localhost:8080`;

const authConfig = {
  hostname: HOST,
  onError: onAuthError,
};

export const authClient = new rpcAuthClient(authConfig);
export const lodasClient = new rpcLodasClient(authConfig);
