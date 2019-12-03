import gql from "graphql-tag";
import { useQuery, useMutation } from '@apollo/react-hooks';

const CONTINUE_DOUBLE_SIGN =
    gql`
    mutation ContinueDoubleSign($continueID: String!, $newHolderID: String!, $newHolderPassword: String!)
    {continueDoubleSign(request: {continueID: $continueID, newHolderID: $newHolderID, newHolderPassword: $newHolderPassword})}
`;

export function ContinueDoubleSign(continueID, newHolderID, newHolderPassword) {
    const res = useMutation(CONTINUE_DOUBLE_SIGN, {variables: {continueID, newHolderID, newHolderPassword}})
    res[0].call()
}
