import gql from "graphql-tag";
import { useQuery, useMutation, ApolloProvider } from '@apollo/react-hooks';


/**
 * Query to list all shipments
 */
const LIST_ALL_SHIPMENTS =
    gql`
    query {listAllShipments {
    id, history{event}
    }}
`;

/**
 * Function executing query for listing all shipments
 */
export function ListAllShipments()  {
    const data = useQuery(LIST_ALL_SHIPMENTS);
    return data
}


/**
 * Query to get shipments details
 * Has the shipment ID as input parameter
 */
const GET_SHIPMENT_DETAILS =
gql`
    query GetShipmentDetails($shipmentID: ID!){getShipmentDetails(shipmentID: $shipmentID) {
    wasteType, id, currentHolderID, producingCompanyID, history{ownerID, event, receiverID}}}
`;


/**
 * Function executing query to get shipment details
 * @param {string} shipmentID
 */
export function GetShipmentDetails(shipmentID)  {
    const data = useQuery(GET_SHIPMENT_DETAILS, { variables: {shipmentID}});
    return data
}

/**
 * Query to create a shipment
 */
const CREATE_SHIPMENT =
gql`
    mutation CreateShipment($wasteType: String!, $currentHolderID: String!, $location: String!, $password: String!)
    {createShipment(request:
        {
            wasteType: $wasteType,
            currentHolderID: $currentHolderID,
            location: $location,
            password: $password
        }
    )}
`;

export function CreateShipment(wasteType, currentHolderID, location, password)  {
    const mutation = useMutation(CREATE_SHIPMENT, { variables: {wasteType, currentHolderID, location, password}});
    mutation[0].call()
}


