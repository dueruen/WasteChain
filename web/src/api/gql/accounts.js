import gql from "graphql-tag";
import { useQuery, useMutation } from '@apollo/react-hooks';

/**
 * Query to list all companies
 */
const LIST_ALL_COMPANIES =
    gql`
    query {listAllCompanies {
    id, name, address{roadName, number, ZipCode }contactinfo{phoneNumber, mail}
    }}
`;

/**
 * Function executing query for listing all companies
 */
export function ListAllCompanies()  {
    const data = useQuery(LIST_ALL_COMPANIES);
    return data
}


/**
 * Query to get company details
 * Has the company ID as input parameter
 */
const GET_COMPANY =
gql`
    query GetCompany($companyID: ID!){getCompany(companyID: $companyID) {
    id, name, address {roadName, number, ZipCode }contactinfo{phoneNumber, mail}, employees{id, name}
    }}
`;


/**
 * Function executing query to get company details
 * @param {string} companyID
 */
export function GetCompany(companyID)  {
    const data = useQuery(GET_COMPANY, { variables: {companyID}});
    return data
}

/**
 * Mutation to create a company
 */
const CREATE_COMPANY =
gql`
    mutation CreateCompany($authData: AuthData!,$name: String!, $address: CreateAddress!, $contactinfo: CreateContactInfo!)
    {createCompany(company:
        {
            authData: $authData,
            name: $name,
            address: $address,
            contactinfo: $contactinfo
        }
    ){id}}
`;

/**
 * Function to create a company
 * @param {AuthData} authData
 * @param {String} name
 * @param {CreateCompany} address
 * @param {CreateContactInfo} contactinfo
 */
export function CreateCompany(authData, name, address, contactinfo)  {
    const res = useMutation(CREATE_COMPANY, { variables: {authData, name, address, contactinfo}});
    res[0].call()
}




/**
 * Mutation to create an employee
 */
const CREATE_EMPLOYEE =
gql`
    mutation CreateEmployee($authData: AuthData!,$name: String!, $companyID: String!)
    {createEmployee(employee:
        {
            authData: $authData,
            name: $name,
            companyID: $companyID,
        }
    ){id}}
`;

/**
 * Function to create employee
 * @param {AuthData} authData
 * @param {string} name
 * @param {string} companyID
 */
export function CreateEmployee(authData, name, companyID)  {
    const res = useMutation(CREATE_EMPLOYEE, { variables: {authData, name, companyID}});
    res[0].call()
}


/**
 * Query to get employee details
 * Has the employee ID as input parameter
 */
const GET_EMPLOYEE =
gql`
    query GetEmployee($employeeID: ID!){getEmployee(employeeID: $employeeID) {
    id, name
    }}
`;


/**
 * Function executing query to get employee details
 * @param {string} employeeID
 */
export function GetEmployee(employeeID)  {
    const data = useQuery(GET_EMPLOYEE, { variables: {employeeID}});
    return data
}



/**
 * Query to list all employees in a company
 */
const LIST_ALL_EMPLOYEES_IN_COMPANY =
    gql`
    query ListAllEmployeesInCompany($companyID: ID!)
    {listAllEmployeesInCompany(companyID: $companyID){id, name}}
`;

/**
 * Function to execute listAllEmployeesInCompany query
 * @param {string} companyID
 */
export function ListAllEmployeesInCompany(companyID)  {
    const data = useQuery(LIST_ALL_EMPLOYEES_IN_COMPANY, { variables: {companyID}});
    return data
}
