import { combineReducers } from "@reduxjs/toolkit";
import common from "./feature/common/slice";
import user from "./feature/user/slice";

const rootReducer = combineReducers({
  common,
  user,
});

export type RootState = ReturnType<typeof rootReducer>;

export default rootReducer;
