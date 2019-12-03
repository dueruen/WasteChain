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
 * Query to create an employee
 */
const CREATE_EMPLOYEE =
gql`
    mutation CreateEmployee($name: String!, $companyID: ID!)
    {createEmployee(employee:
        {
            name: $name,
            companyID: $companyID,
        }
    )}
`;

export function CreateEmployee(name, companyID)  {
    const res = useMutation(CREATE_EMPLOYEE, { variables: {name, companyID}});
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
 * Function executing query for listing all employees in a company
 */
export function ListAllEmployeesInCompany(companyID)  {
    const data = useQuery(LIST_ALL_EMPLOYEES_IN_COMPANY, { variables: {companyID}});
    return data
}
