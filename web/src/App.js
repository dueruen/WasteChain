import React, { Fragment } from 'react';
// import './App.css';

import { ApolloClient } from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloProvider, useQuery } from '@apollo/react-hooks';
import { HttpLink } from 'apollo-link-http';

import Pages from './pages';

const cache = new InMemoryCache();
const link = new HttpLink({
  uri: 'http://localhost:8081'
});

const client = new ApolloClient({
  cache,
  link
});

const App = () => {
    return (
        <ApolloProvider client={client}>
            {/* <IsLoggedIn /> */}
            <Pages></Pages>
        </ApolloProvider>
    )
}

export default App;
