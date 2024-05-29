import { env } from "$lib";
import type { CookieInfo } from "$lib/helpers/helpers";

export async function githubLink(cookie: CookieInfo): Promise<string> {
    const res = await fetch(`${env.VITE_URL}/api/v1/auth/github/link`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const r = await res.json();
    console.log(r);
    return r["url"]
}


export async function githubLinkCallback(cookie: CookieInfo, options: {
    installationId: string,
    code: string,
    state: string,
}): Promise<string> {
    if (options.code.length === 0) {
        return ""
    }

    const res = await fetch(`${env.VITE_URL}/api/v1/auth/github/callback?code=${options.code}&installation_id=${options.installationId}&state=${options.state}`, {
        method: "POST",
        credentials: 'include',
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const r = await res.json();

    return "success"
}

export async function getAvailableRepos(cookie: CookieInfo, search?: string): Promise<GithubRepository[]> {
    const res = await fetch(`${env.VITE_URL}/api/v1/cloud/repos`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const methods: GithubRepository[] = await res.json();
    return methods
}