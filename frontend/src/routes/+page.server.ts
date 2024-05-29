import { env } from "$lib";
import { redirect } from "@sveltejs/kit";
import {
  githubCallback,
  githubLogin,
  googleCallback,
  googleLogin,
} from "$lib/services/auth/services.js";

export const load = async ({ cookies, url }) => {
  let error: string | undefined = undefined;
  let params = url.searchParams.size;
  if (params !== 0) {
    if (url.searchParams.get("state")?.includes("github")) {
      const res = await githubCallback(
        String(url.searchParams.get("code")),
        String(url.searchParams.get("state")),
        "signin"
      );
      console.log(res);

      if (res.token !== undefined) {
        cookies.set("token", res.token, {
          path: "/",
          httpOnly: true,
          sameSite: "strict",
          secure: Boolean(env.VITE_PRODUCTION),
          maxAge: 60 * 60 * 24 * 30,
        });
        throw redirect(303, "/stores");
      }
    }

    if (url.searchParams.get("state")?.includes("google")) {
      const res = await googleCallback(
        String(url.searchParams.get("code")),
        String(url.searchParams.get("state")),
        "signin"
      );
      console.log(res);

      if (res.token !== undefined) {
        cookies.set("token", res.token, {
          path: "/",
          httpOnly: true,
          sameSite: "strict",
          secure: Boolean(env.VITE_PRODUCTION),
          maxAge: 60 * 60 * 24 * 30,
        });
        throw redirect(303, "/stores");
      }
    }
  }

  return {
    err: error,
  };
};

export const actions = {
  login: async ({ request, cookies }) => {
    const data = await request.formData();

    const email = data.get("email");
    const password = data.get("password");

    const res = await fetch(`${env.VITE_URL}/api/auth/email/auth`, {
      method: "POST",
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    });

    const resData = await res.json();
    if (res.status === 200) {
      console.log(resData);

      cookies.set("token", resData.token, {
        path: "/",
        httpOnly: true,
        sameSite: "strict",
        secure: Boolean(env.VITE_PRODUCTION),
        maxAge: 60 * 60 * 24 * 30,
      });
      throw redirect(303, "/profile");
    }

    return {
      error: resData.error,
    }
  },

  githubLogin: async () => {
    const url = await githubLogin();
    throw redirect(301, url);
  },

  googleLogin: async () => {
    const url = await googleLogin();
    throw redirect(303, url);
  },
};
