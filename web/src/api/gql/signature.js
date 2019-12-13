import gql from "graphql-tag";
import { useMutation } from '@apollo/react-hooks';

/**
 * Mutation to continue double sign process
 */
const CONTINUE_DOUBLE_SIGN =
    gql`
    mutation ContinueDoubleSign($continueID: String!, $newHolderID: String!, $newHolderPassword: String!)
    {continueDoubleSign(request: {continueID: $continueID, newHolderID: $newHolderID, newHolderPassword: $newHolderPassword})}
`;

/**
 * Function to continue double sign process
 * @param {String} continueID
 * @param {String} newHolderID
 * @param {String} newHolderPassword
 */
export function ContinueDoubleSign(continueID, newHolderID, newHolderPassword) {
    const res = useMutation(CONTINUE_DOUBLE_SIGN, {variables: {continueID, newHolderID, newHolderPassword}})
    res[0].call()
}
