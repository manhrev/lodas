import { combineReducers } from "@reduxjs/toolkit";
import common from "./feature/common/slice";
import user from "./feature/user/slice";
import sheet from "./feature/sheet/slice";

const rootReducer = combineReducers({
  common,
  user,
  sheet,
});

export type RootState = ReturnType<typeof rootReducer>;

export default rootReducer;
