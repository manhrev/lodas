import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";
import moment from "moment";

export const formatDateTime = (time: Timestamp.AsObject) => {
  const date = moment.unix(time.seconds).format("DD/MM/YYYY HH:mm");

  return date;
};

export const formatDate = (time: Timestamp.AsObject) => {
  const date = moment.unix(time.seconds).format("DD/MM/YYYY");
  return date;
};
