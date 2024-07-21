'use client'

import { useState, useEffect } from 'react';
import { BACKEND_URI } from '@/lib/consts';
import RSVPForm from '@/components/RSVPForm';
import { useRouter } from 'next/navigation';
import { useToast } from '@/components/ui/use-toast';

export const runtime = 'edge';

export default function RSVPPage({ params }: { params: { slug: string } }) {
    const router = useRouter();
    const { toast } = useToast();

    const [loading, setLoading] = useState<boolean>(true);
    const [name, setName] = useState<string>('');

    useEffect(() => {
        const setup = async () => {
            const status_response = await fetch(BACKEND_URI + "/status/" + params.slug, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
            });

            if (!status_response.ok) {
                toast({
                    title: "Uh oh! Something went wrong.",
                    description: "There was a problem with your request.",
                });
            }

            const status: boolean = await status_response.json();
            if (status) {
                router.push('/thank-you')
            }

            const name_response = await fetch(BACKEND_URI + "/name/" + params.slug, {
                method: "GET",
                mode: "cors",
                headers: { "Content-Type": "application/json" },
            });

            if (!name_response.ok) {
                toast({
                    title: "Uh oh! Something went wrong.",
                    description: "There was a problem with your request.",
                });
            }

            const name: string = await name_response.json();
            console.log(name)
            setName(name);
            setLoading(false);
        }

        setup();
    }, [])

    
    if (loading) {
        return (
            <></>
        );
    }

    return (
        <RSVPForm id={params.slug} name={name}/>
    );
}
