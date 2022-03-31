import { selectToken } from './../../store/index';
import { useSelector } from 'react-redux';

import {GraphQLClient} from 'graphql-request'

const graphqlReq = new GraphQLClient(`http://${process.env.NEXT_PUBLIC_API_URL}/query`)

export const graphqlRequest = () => {
    const currentToken = useSelector(selectToken)
    if (currentToken === null) {
        return graphqlReq
    } else {
        return  graphqlReq.setHeader("authorization", currentToken!);
    }
}

