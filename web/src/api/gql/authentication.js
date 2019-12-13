import gql from "graphql-tag";
import { useMutation } from '@apollo/react-hooks';

/**
 * Mutation to log in
 */
const LOGIN =
    gql`
    mutation Login($username: String!, $password: String!)
    {login(request: {username: $username, password: $password})}
`;

/**
 * Function to log in
 * @param {String} username
 * @param {String} password
 */
export function Login(username, password) {
    const res = useMutation(LOGIN, {variables: {username, password}})
    res[0].call()
}
