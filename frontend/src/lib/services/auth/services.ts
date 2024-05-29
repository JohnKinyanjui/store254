import { env } from "$lib";


export async function githubLogin(): Promise<string> {
    const res = await fetch(`${env.VITE_URL}/api/auth/oauth2/github/app`, {
        method: "POST",
    })

    const r = await res.json();

    return r["url"]
}

export async function githubCallback(code: string, state: string,type: string): Promise<any> {
    if (code.length === 0) {
        return ""
    }
    const res = await fetch(`${env.VITE_URL}/api/auth/oauth2/github/app/callback?code=${code}&state=${state}&type=${type}`, {
        method: "POST",
        credentials: 'include',
    })

    const r = await res.json();

    return r
}


export async function googleLogin(): Promise<string> {
    const res = await fetch(`${env.VITE_URL}/api/auth/oauth2/goole`, {
        method: "POST",
    })

    const r = await res.json();

    return r["url"]
}

export async function googleCallback(code: string, state: string, type: string): Promise<any> {
    if (code.length === 0) {
        return ""
    }
    const res = await fetch(`${env.VITE_URL}/api/auth/oauth2/github/app/callback?code=${code}&state=${state}&type=${type}`, {
        method: "POST",
        credentials: 'include',
    })

    const r = await res.json();

    return r
}
