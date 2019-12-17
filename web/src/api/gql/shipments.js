import gql from "graphql-tag";
import { useQuery, useMutation } from '@apollo/react-hooks';

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
 * Query to get shipments details
 * Has the shipment ID as input parameter
 */
const GET_SHIPMENT_DETAILS =
gql`
    query GetShipmentDetails($shipmentID: ID!){getShipmentDetails(shipmentID: $shipmentID) {
    wasteType, id, currentHolderID, producingCompanyID, history{ownerID, event, receiverID}}}
`;


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


