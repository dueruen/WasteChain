import React, { Fragment } from 'react';

import { CreateAccountForm } from '../components';
import { useApolloClient, useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';

export const CREATE_COMPANY = gql`
  mutation createCompany($name: String!) {
    createCompany(name: $name)
  }
`;

const CreateAccountPage = () => {
    const client = useApolloClient();
    const [create, { loading, error }] = useMutation(
        CREATE_COMPANY,
        {
            onCompleted( {create} ) {
                console.log('I was here!!!')
            }
        }
    )

    if (loading) return <div>LOADING...</div>
    if (error) return (
        console.log(error),
        <div>ERROR!!!</div>
    );

    return <CreateAccountForm create={create}/>;
}

export default CreateAccountPage;
