import React, { Fragment } from 'react';
import { useQuery } from '@apollo/react-hooks';
import gql from 'graphql-tag';

//import { LaunchTile, Header, Button, Loading } from '../components';

const ls = ({ companies }) => {
    const { id, name } = companies;
    return (
      <div>
          id: {companies.id}
      </div>
    );
  };

export const LAUNCH_TILE_DATA = gql`
  fragment ls on Company {
    id
    name
  }
`;

export const GET_COMPANIES = gql`
  query listAllCompanies($after: String) {
    companies(after: $after) {
      id,
      name
    }
  }
  ${LAUNCH_TILE_DATA}
`;

export default function Companies() {
  const { data, loading, error, fetchMore } = useQuery(GET_COMPANIES);
  if (loading) return <div>LOADING</div>;
  if (error) return <p>ERROR</p>, console.log("ERROR: " + error), console.log("DATA: " + data);

  console.log(data);
  return (
    <Fragment>
      <div />
      <h1>Companies</h1>
      {/* {data.launches &&
        data.launches.launches &&
        data.launches.launches.map(launch => (
          <LaunchTile key={launch.id} launch={launch} />
        ))} */}
      {/* {data.launches &&
        data.launches.hasMore && (
          <Button
            onClick={() =>
              fetchMore({
                variables: {
                  after: data.launches.cursor,
                },
                updateQuery: (prev, { fetchMoreResult, ...rest }) => {
                  if (!fetchMoreResult) return prev;
                  return {
                    ...fetchMoreResult,
                    launches: {
                      ...fetchMoreResult.launches,
                      launches: [
                        ...prev.launches.launches,
                        ...fetchMoreResult.launches.launches,
                      ],
                    },
                  };
                },
              })
            }
          >
            Load More
          </Button>
        )} */}
    </Fragment>
  );
}
