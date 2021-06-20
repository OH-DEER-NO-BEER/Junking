using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;
using static Hands_type;

public class TimerCtrl : MonoBehaviour
{
    public AudioClip audioClip1;
    private AudioSource audioSource;

    Slider slider;
    float time = 0;

    [HideInInspector]
    public bool selected = false;

    [HideInInspector]
    public Hands hand = Hands.Paper;

    private int Random_Select_Hands()
    {
        return Random.Range((int)Hands.Rock, (int)(Hands.Paper + 1));
    }

    // Start is called before the first frame update
    void Start()
    {
        slider = GetComponent<Slider>();
        time = slider.maxValue;
        audioSource = gameObject.GetComponent<AudioSource>();
    }

    // Update is called once per frame
    void Update()
    {
        if (time == 20.0f)
        {
            audioSource.Play();
        }

        time -= Time.deltaTime;

        if(time < 5.0f)
        {
            audioSource.Stop();
            audioSource.clip = audioClip1;
            audioSource.Play();
        }

        if (time < slider.minValue)
        {
            hand = Trans_Hand(Random_Select_Hands());
            selected = true;
        }

        slider.value = time;
    }
}
