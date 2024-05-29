import { env } from "$lib";
import { redirect } from "@sveltejs/kit";

export async function load({ url }) {
    const email = String(url.searchParams.get("token"));
    if (email === "null" || email?.length === 0) {

        throw redirect(303, "/");
    }

    return {
        token: email
    }
}

export const actions = {
    verification: async ({ cookies, request }) => {
        const data = await request.formData()
        const token = String(data.get("token"))
        const verificationCode = String(data.get("verification_code"))

        if (verificationCode.length === 0) {
            return { error: "make sure your verification code is empty" }
        }

        const res = await fetch(`${env.VITE_URL}/api/auth/email/validate`, {
            method: "POST",
            body: JSON.stringify({
                otp_id: token,
                otp: verificationCode,
            })
        })

        if (res.status === 200) {
            const resData = await res.json();

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


        return { error: "make sure your verification code is correct" }

    }
};