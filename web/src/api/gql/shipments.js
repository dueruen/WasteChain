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
    )ID}
`;

/**
 * Creates a new shipment
 * @param {string} wasteType
 * @param {string} currentHolderID
 * @param {string} location
 * @param {string} password
 */
export function CreateShipment(wasteType, currentHolderID, location, password)  {
    const mutation = useMutation(CREATE_SHIPMENT, { variables: {wasteType, currentHolderID, location, password}});
    mutation[0].call()
}

/**
 * Mutation to transfer a shipment
 */
const TRANSFER_SHIPMENT =
gql`
    mutation TransferShipment($shipmentID: String!, $ownerID: String!,$receiverID: String!, $location: String!, $password: String!)
    {tansferShipment(request:
        {
            shipmentID: $shipmentID,
            ownerID: $ownerID,
            receiverID: $receiverID,
            location: $location,
            password: $password,
        }
    )}
`;

/**
 * Mutation function to transfer a shipment
 * @param {String} shipmentID
 * @param {String} ownerID
 * @param {String} receiverID
 * @param {String} location
 * @param {String} password
 */
export function TransferShipment(shipmentID, ownerID, receiverID, location, password)  {
    const mutation = useMutation(TRANSFER_SHIPMENT, { variables: {shipmentID, ownerID, receiverID, location, password}});
    mutation[0].call()
}


/**
 * Mutation to process a shipment
 */
const PROCESS_SHIPMENT =
gql`
    mutation ProcessShipment($shipmentID: String!, $ownerID: String!, $location: String!, $password: String!)
    {processShipment(request:
        {
            shipmentID: $shipmentID,
            ownerID: $ownerID,
            location: $location,
            password: $password,
        }
    )}
`;

/**
 * Mutation function to process a shipment
 * @param {String} shipmentID
 * @param {String} ownerID
 * @param {String} location
 * @param {String} password
 */
export function ProcessShipment(shipmentID, ownerID, location, password)  {
    const mutation = useMutation(PROCESS_SHIPMENT, { variables: {shipmentID, ownerID, location, password}});
    mutation[0].call()
}

