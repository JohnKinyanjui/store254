import { env, validateEmail, validatePassword } from "$lib";
import { redirect } from "@sveltejs/kit";

/** @type {import('./$types').Actions} */
export const actions = {
    sign_up: async ({ request }) => {
        const data = await request.formData()


        const name = String(data.get("full_name"))
        const email = String(data.get("email"))
        const password = String(data.get("password"))


        if (name.length < 5) {
            return { error_name: "make sure the full name has at least 5 characters" }
        }
        if (!validateEmail(email)) {
            return { error_email: "make sure the email is valid" }
        }

        const err = validatePassword(password, String(data.get("re-password")))
        if (err !== null) {
            return { error_password: err }
        }

        const response = await registerUserWithEmail({
            name: name,
            email: email,
            password: password
        })

        if (response.status === 200) {
            var res = await response.json()

            throw redirect(303, `/register/verify?token=${res.id}`);
        }
    }
};

async function registerUserWithEmail(data: { name: string | null, email: string, password: string }) {
    const res = await fetch(`${env.VITE_URL}/api/auth/email/register`, {
        method: "POST",
        body: JSON.stringify({
            full_name: data.name,
            email: data.email,
            password: data.password,
        })
    })

    return res
}