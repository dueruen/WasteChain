import { ApolloClient } from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { HttpLink } from 'apollo-link-http';
import { setContext } from 'apollo-link-context';

const cache = new InMemoryCache();
console.log("------------Start")
console.log(process.env);
console.log("------------END")

let host = "localhost";
if (process.env.REACT_APP_HOST) {
    host = process.env.REACT_APP_HOST
}

const link = new HttpLink({
  uri: 'http://' + host + ':8081/query',
  credentials: 'same-origin'
})

const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem('token');
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const client = new ApolloClient({
  link: authLink.concat(link),
  cache
});

export default client

