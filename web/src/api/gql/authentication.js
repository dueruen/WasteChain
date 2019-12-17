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


