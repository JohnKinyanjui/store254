// place files you want to import through the `$lib` alias in this folder.
import * as validators from './helpers/validators'
import * as config from './helpers/config'
import { writable } from 'svelte/store';
import * as customerService from './services/customers/services'
import * as productService from './services/products/services'



export const env = import.meta.env

export const loadingPage = writable(false);

// views

// validators
export const validateEmail = validators.validateEmail
export const validateCategoryForm = validators.validateCategoryForm
export const validateProductForm = validators.validateProductForm
export const validatePassword = validators.validatePassword
export const navItems = config.navItems

//


// services

// customers
export const getCustomers = customerService.getCustomers;


//products
export const getProducts = productService.getProducts;