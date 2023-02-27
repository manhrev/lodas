import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";

export const formatDateTime = (time: Timestamp.AsObject) => {
  const date = new Date(time.seconds * 1000);
  const yyyy = date.getFullYear();
  let mm = date.getMonth() + 1; // Months start at 0!
  let dd = date.getDate();
  let mmstring = mm.toString();
  let ddstring = dd.toString();
  if (dd < 10) ddstring = "0" + dd;
  if (mm < 10) mmstring = "0" + mm;

  return dd + "/" + mm + "/" + yyyy + " " + date.toLocaleTimeString();
};

export const formatDate = (time: Timestamp.AsObject) => {
  const date = new Date(time.seconds * 1000);
  const yyyy = date.getFullYear();
  let mm = date.getMonth() + 1; // Months start at 0!
  let dd = date.getDate();
  let mmstring = mm.toString();
  let ddstring = dd.toString();
  if (dd < 10) ddstring = "0" + dd;
  if (mm < 10) mmstring = "0" + mm;

  return dd + "/" + mm + "/" + yyyy;
};
