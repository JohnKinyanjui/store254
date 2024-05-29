import { env } from "$lib";
import type { Cookies } from "@sveltejs/kit";

export const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];


export function formatDateToLocal(timestamp: string): string {
    if (timestamp == "0001-01-01T00:00:00Z") {
        return "Not Available";
    }
    const utcDateTime = new Date(timestamp);

    // Format the date and time in local time
    const options: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: true,
    };

    return utcDateTime.toLocaleString(undefined, options).toUpperCase();
}

export function formatInputTimestampToLocal(timestamp: string): string {
    const utcDateTime = new Date(timestamp);

    // Format the date and time in local time for "datetime-local" input
    const year = utcDateTime.getFullYear();
    const month = String(utcDateTime.getMonth() + 1).padStart(2, '0');
    const day = String(utcDateTime.getDate()).padStart(2, '0');
    const hours = String(utcDateTime.getHours()).padStart(2, '0');
    const minutes = String(utcDateTime.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day}T${hours}:${minutes}`;
}

export function formatToUTC(datetimeLocalValue: string) {
    if (datetimeLocalValue.length == 0) {
        return "0001-01-01T00:00:00Z"
    }

    // Create a Date object from the input value
    const localDate = new Date(datetimeLocalValue);

    // Get the UTC date and time components
    const year = localDate.getUTCFullYear();
    const month = String(localDate.getUTCMonth() + 1).padStart(2, '0'); // Months are zero-based
    const day = String(localDate.getUTCDate()).padStart(2, '0');
    const hours = String(localDate.getUTCHours()).padStart(2, '0');
    const minutes = String(localDate.getUTCMinutes()).padStart(2, '0');

    // Format the UTC date and time as a string
    const utcString = `${year}-${month}-${day}T${hours}:${minutes}:00Z`;

    return utcString;
}


export function getTimeDifference(utcTimestamp: string): string {
    const now = new Date();
    const pastDate = new Date(utcTimestamp);
    const timeDifferenceInSeconds = Math.floor((now.getTime() - pastDate.getTime()) / 1000);

    if (timeDifferenceInSeconds < 60) {
        return `${timeDifferenceInSeconds} second${timeDifferenceInSeconds === 1 ? '' : 's'} ago`;
    } else if (timeDifferenceInSeconds < 3600) {
        const minutes = Math.floor(timeDifferenceInSeconds / 60);
        return `${minutes} minute${minutes === 1 ? '' : 's'} ago`;
    } else if (timeDifferenceInSeconds < 86400) {
        const hours = Math.floor(timeDifferenceInSeconds / 3600);
        return `${hours} hour${hours === 1 ? '' : 's'} ago`;
    } else if (timeDifferenceInSeconds < 604800) {
        const days = Math.floor(timeDifferenceInSeconds / 86400);
        return `${days} day${days === 1 ? '' : 's'} ago`;
    } else if (timeDifferenceInSeconds < 2419200) {
        const weeks = Math.floor(timeDifferenceInSeconds / 604800);
        return `${weeks} week${weeks === 1 ? '' : 's'} ago`;
    } else {
        const months = Math.floor(timeDifferenceInSeconds / 2419200);
        return `${months} month${months === 1 ? '' : 's'} ago`;
    }
}

// // Example usage
// const utcTimestamp = "2023-10-10T12:50:57.413846Z";
// const timeDifference = getTimeDifference(utcTimestamp);
// console.log(`Time Difference: ${timeDifference}`);


export interface CookieInfo {
    token: string,
    current_store: string
}

export function getCookieInfo(cookies: Cookies): CookieInfo {
    return {
        token: String(cookies.get("token") ?? "",),
        current_store: String(cookies.get("current_store") ?? ""),
    }
}

export function formatNumberWithCommas(number: number): string {
    // Check if the input is a valid number
    if (isNaN(number)) {
        return "Invalid Number";
    }

    // Convert the number to a string
    const numberString = number.toString();

    // Split the string into integer and decimal parts
    const parts = numberString.split(".");

    // Format the integer part with commas
    const integerPart = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ",");

    // Reconstruct the formatted number
    if (parts.length === 2) {
        return `${integerPart}.${parts[1]}`;
    } else {
        return integerPart;
    }
}

export function getGreeting(): string {
    const currentHour = new Date().getHours();
    if (currentHour < 12) {
        return 'Good morning';
    } else if (currentHour < 18) {
        return 'Good afternoon';
    } else {
        return 'Good evening';
    }
}
