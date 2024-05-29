
export function validateProductForm(name: string | undefined, description: string | undefined, regularPrice: string | undefined, salesPrice: string | undefined): Record<string, any> | null {
    if (name == null || name.length == 0) {
        return {
            error: "please enter your product name",
            error_product_name: "please enter your product name"
        }
    }


    if (description == null || description.length == 0) {
        return {
            error: "please enter your product description",
            error_product_description: "please enter your product description"
        }
    }

    if (regularPrice == null || regularPrice.length == 0) {
        return {
            error: "please enter your product regular price",
            error_product_regular_price: "please enter your product regular price"
        }
    }

    if (regularPrice != null && salesPrice != null) {
        if (Number(salesPrice) > Number(regularPrice)) {
            return {
                error: "sales price cannot be greater that regular price",
                error_product_sales_price: "sales price cannot be greater that regular price"
            }
        }
    }

    return null
}

export function validateCategoryForm(name: string | undefined) {
    if (name == null || name.length == 0) {
        return {
            error_category_name: "please enter your product name"
        }
    }

    return null
}

export function validateEmail(email: string): boolean {
    const expression: RegExp = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;
    return expression.test(email);
}

export function validatePassword(password: string, repassword: string): string | null {
    // Define the criteria for a strong password
    const minLength = 8;
    const hasUpperCase = /[A-Z]/.test(password);
    const hasLowerCase = /[a-z]/.test(password);
    const hasNumbers = /\d/.test(password);
    const hasSpecialChars = /[!@#$%^&*()-_=+{};:',.<>?/\\|[\]~]/.test(password);

    // Check if the password meets the criteria
    if (
        password.length < minLength ||
        !hasUpperCase ||
        !hasLowerCase ||
        !hasNumbers ||
        !hasSpecialChars
    ) {
        // Return a string describing the reason why the password is not strong
        return "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character.";
    }

    // Check if the password matches the "repassword" parameter
    if (password !== repassword) {
        return "Passwords do not match.";
    }

    // Return null if both conditions are met, indicating a strong password
    return null;
}