import gql from "graphql-tag";

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
 * Query to list all employees in a company
 */
const LIST_ALL_EMPLOYEES_IN_COMPANY =
    gql`
    query ListAllEmployeesInCompany($companyID: ID!)
    {listAllEmployeesInCompany(companyID: $companyID){id, name}}
`;
