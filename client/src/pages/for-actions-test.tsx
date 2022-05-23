import Link from 'next/link'
import { GetServerSideProps } from 'next';

const ForActionsTest = () => {
    return (
        <div>
            Enter
        </div>
    );
}


export const getServerSideProps:GetServerSideProps = async (ctx) => {


    return {
        props:{
            data:null
        }
    }
}

export default ForActionsTest;
