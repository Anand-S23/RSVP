'use client'


import { BACKEND_URI } from '@/lib/consts';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';
import { useToast } from '../ui/use-toast';

interface FormProps {
    id: string,
    name: string
}

const RSVPForm: React.FC<FormProps> = (props: FormProps) => {
    const router = useRouter();
    const { toast } = useToast();

    const [isAttending, setIsAttending] = useState<string>('yes');
    const [arrivingDate, setArrivingDate] = useState<string>('');
    const [leavingDate, setLeavingDate] = useState<string>('');
    const [comments, setComments] = useState<string>('');
    const [error, setError] = useState<string | null>(null);

    const handleAttendingChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setIsAttending(event.target.value);
        setError(null); // Clear the error when the attending status changes
    };

    const handleArrivingDateChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setArrivingDate(event.target.value);
    };

    const handleLeavingDateChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLeavingDate(event.target.value);
    };

    const handleCommentsChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setComments(event.target.value);
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (isAttending === 'yes' && (!arrivingDate || !leavingDate)) {
            setError('Please provide both arriving and leaving dates.');
            return;
        }

        setError(null); // Clear any existing errors if validation passes

        const submit_data = {
            "id": props.id,
            "name": props.name,
            "answered": true,
            "attending": isAttending === 'yes',
            "date_arriving": new Date(arrivingDate),
            "date_departure": new Date(leavingDate),
            "comments": comments
        }

        console.log(submit_data);

         const rsvp_response = await fetch(BACKEND_URI + '/rsvp', {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(submit_data),
            headers: { "Content-Type": "application/json" },
        });

        if (!rsvp_response.ok) {
            toast({
                title: "Uh oh! Something went wrong.",
                description: "There was a problem with your request.",
            });
        }

        const rsvp = await rsvp_response.json();
        console.log(rsvp)
        router.push('/thank-you')
    };

    return (
        <div className="flex justify-center items-center min-h-screen bg-gray-100">
            <div className="bg-white p-8 shadow-lg w-full max-w-lg">
                <h2 className="text-2xl font-semibold text-center">Gurpreet & Satwinder</h2>
                <h3 className="text-lg mb-4 text-center">Invite you, {props.name}, to share their wedding on September 7th 2024</h3>

                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="attending">
                            Will you be attending?
                        </label>
                        <select
                            className="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            id="attending"
                            value={isAttending}
                            onChange={handleAttendingChange}
                        >
                            <option value="yes">Yes</option>
                            <option value="no">No</option>
                        </select>
                    </div>
                    
                    {isAttending === 'yes' && (
                        <>
                            <div className="mb-4">
                                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="arriving">
                                    Date Arriving
                                </label>
                                <input
                                    className="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                                    id="arriving"
                                    type="date"
                                    value={arrivingDate}
                                    onChange={handleArrivingDateChange}
                                />
                            </div>
                            <div className="mb-4">
                                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="leaving">
                                    Date Leaving
                                </label>
                                <input
                                    className="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                                    id="leaving"
                                    type="date"
                                    value={leavingDate}
                                    onChange={handleLeavingDateChange}
                                />
                            </div>
                        </>
                    )}

                    <div className="mb-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="comments">
                            Comments
                        </label>
                        <textarea
                            className="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            id="comments"
                            placeholder="Any comments or special requirements"
                            value={comments}
                            onChange={handleCommentsChange}
                        />
                    </div>

                    {error && <div className="mb-4 text-red-500">{error}</div>}

                    <div className="flex items-center justify-between">
                        <button
                            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 focus:outline-none focus:shadow-outline"
                            type="submit"
                        >
                            Submit
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default RSVPForm;

