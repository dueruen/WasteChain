import gql from "graphql-tag";
import { useQuery, ApolloProvider } from '@apollo/react-hooks';

const LIST_ALL_SHIPMENTS =
    gql`
    query {listAllShipments {
    id, history{event}
    }}`
;

export function ListAllShipments()  {
    const data = useQuery(LIST_ALL_SHIPMENTS);
    return data
}



