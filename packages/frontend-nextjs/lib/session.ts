import { SessionOptions } from "iron-session";

export interface SessionData {
  userId?: string;
  email?: string;
  name?: string;
  accessToken?: string;
}

export const defaultSession: SessionData = {};

export const sessionOptions: SessionOptions = {
  password:
    process.env.COOKIE_PASSWORD || "Yzf76MDXgiLKUvKUWacsPvwVs6VLC2aLUYmTwG",
  cookieName: "dialog-session",
  cookieOptions: {
    httpOnly: true,
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
  },
  ttl: 4200,
};
